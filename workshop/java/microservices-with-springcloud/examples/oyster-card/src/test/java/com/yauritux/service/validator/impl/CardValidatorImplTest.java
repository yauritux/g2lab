package com.yauritux.service.validator.impl;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.when;

import java.math.BigDecimal;
import java.util.Optional;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.exception.CardException;
import com.yauritux.exception.CardNotFoundException;
import com.yauritux.exception.InvalidCardOperation;
import com.yauritux.model.constant.CardOperationType;
import com.yauritux.service.query.CardQueryService;
import com.yauritux.service.validator.CardValidator;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
public class CardValidatorImplTest {
	
	private static final String CARD_SN = "1234567890";

	private CardValidator cardValidator;
	
	@MockBean
	private CardQueryService cardQueryService;
	
	@Before
	public void setup() {
		this.cardValidator = new CardValidatorImpl(cardQueryService);
	}
	
	@Test
	public void validateOwner_nullOwnerName_returnsInvalidCardOperationException() {
		CardException exception = cardValidator.validateOwner(null);
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validateOwner_emptyOwnerName_returnsInvalidCardOperationException() {
		CardException exception = cardValidator.validateOwner(" ");
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validateOwner_ownerNameIsProvided_returnsNull() {
		assertNull(cardValidator.validateOwner("Yauri"));
	}
	
	@Test
	public void validateSerialNumber_nullSerialNo_returnsCardNotFoundException() {
		CardException exception = cardValidator.validateSerialNumber(null);
		assertTrue(exception instanceof CardNotFoundException);
	}
	
	@Test
	public void validateSerialNumber_nullSerialNo_returnsAppropriateMessage() {
		CardException exception = cardValidator.validateSerialNumber(null);
		assertThat(exception.getMessage()).isEqualTo("No serial number found !");
	}
	
	@Test
	public void validateSerialNumber_emptySerialNo_returnsCardNotFoundException() {
		CardException exception = cardValidator.validateSerialNumber(" ");
		assertTrue(exception instanceof CardNotFoundException);
	}	
	
	@Test
	public void validateSerialNumber_emptySerialNo_returnsAppropriateMessage() {
		CardException exception = cardValidator.validateSerialNumber(" ");
		assertThat(exception.getMessage()).isEqualTo("No serial number found !");
	}
	
	@Test
	public void validateSerialNumber_nonExistingSerialNumber_returnsCardNotFoundException() {
		when(cardQueryService.findBySerialNo(CARD_SN)).thenReturn(Optional.empty());
		CardException exception = cardValidator.validateSerialNumber(CARD_SN);
		assertTrue(exception instanceof CardNotFoundException);
	}
	
	@Test
	public void validateSerialNumber_nonExistingSerialNumber_returnsAppropriateMessage() {
		when(cardQueryService.findBySerialNo(CARD_SN)).thenReturn(Optional.empty());
		CardException exception = cardValidator.validateSerialNumber(CARD_SN);
		assertThat(exception.getMessage()).isEqualTo("Cannot find card with serial number '" + CARD_SN + "' !");
	}
	
	@Test
	public void validateAmount_addNegativeAmount_returnsInvalidCardOperationException() {
		CardException exception = cardValidator.validateAmount(BigDecimal.valueOf(-200), CardOperationType.ADD_BALANCE);
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validateAmount_addZeroAmount_returnsInvalidCardOperationException() {
		CardException exception = cardValidator.validateAmount(BigDecimal.ZERO, CardOperationType.ADD_BALANCE);
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validateAmount_addBalanceEverythingOK_returnsNull() {
		assertNull(cardValidator.validateAmount(BigDecimal.valueOf(100), CardOperationType.ADD_BALANCE));
	}
	
	@Test
	public void validateAmount_deductNegativeAmount_returnsInvalidCardOperationException() {
		CardException exception = cardValidator.validateAmount(BigDecimal.valueOf(-100), CardOperationType.DEDUCT_BALANCE);
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validateAmount_deductZeroAmount_returnsInvalidCardOperationException() {
		CardException exception = cardValidator.validateAmount(BigDecimal.ZERO, CardOperationType.DEDUCT_BALANCE);
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validateAmount_deductBalanceEverythingOK_returnsNull() {
		assertNull(cardValidator.validateAmount(BigDecimal.valueOf(50), CardOperationType.DEDUCT_BALANCE));
	}
	
	@Test
	public void validateAmount_initializeCardWithZeroAmount_returnsInvalidCardOperation() {
		CardException exception = cardValidator.validateAmount(BigDecimal.ZERO, CardOperationType.CARD_INITIALIZATION);
		assertTrue(exception instanceof InvalidCardOperation);
	}
	
	@Test
	public void validatAmount_initializeCardEverythingOK_returnsNull() {
		assertNull(cardValidator.validateAmount(BigDecimal.valueOf(30), CardOperationType.CARD_INITIALIZATION));
	}
}

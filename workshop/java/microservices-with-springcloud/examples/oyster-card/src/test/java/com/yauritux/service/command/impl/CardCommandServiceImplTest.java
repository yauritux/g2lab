package com.yauritux.service.command.impl;

import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.when;

import java.math.BigDecimal;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.exception.CardNotFoundException;
import com.yauritux.exception.InvalidCardOperation;
import com.yauritux.model.constant.CardExceptionType;
import com.yauritux.model.constant.CardOperationType;
import com.yauritux.model.entity.Card;
import com.yauritux.repository.CardRepository;
import com.yauritux.service.command.CardCommandService;
import com.yauritux.service.validator.CardValidator;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
public class CardCommandServiceImplTest {
	
	private static final String CARD_SN = "1234567890";
	
	@MockBean
	private CardRepository cardRepository;
	
	@MockBean
	private CardValidator cardValidator;

	private CardCommandService cardCommandService;
	
	@Before
	public void setup() {
		this.cardCommandService = new CardCommandServiceImpl(cardRepository, cardValidator);
	}
	
	@Test(expected = InvalidCardOperation.class)
	public void initializeCard_nullOwnerName_throwInvalidCardOperation() throws Exception {
		when(cardValidator.validateOwner(null)).thenReturn(new InvalidCardOperation("Missing owner name!", CardExceptionType.CARD_MISSING_ATTRIBUTE_EXCEPTION));
		cardCommandService.initializeCard(null, BigDecimal.ZERO);
	}
	
	@Test(expected = InvalidCardOperation.class)
	public void initializeCard_emptyBalance_throwInvalidCardOperation() throws Exception {
		when(cardValidator.validateOwner("Yauri")).thenReturn(null);
		when(cardValidator.validateAmount(BigDecimal.ZERO, CardOperationType.CARD_INITIALIZATION)).thenReturn(new InvalidCardOperation("No amount to be initialized!", CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION));
		cardCommandService.initializeCard("Yauri", BigDecimal.ZERO);
	}
	
	
	@Test(expected = CardNotFoundException.class)
	public void deduct_noSerialNumber_throwCardNotFoundException() throws Exception {
		when(cardValidator.validateSerialNumber(null)).thenReturn(new CardNotFoundException("Cannot find card!"));
		cardCommandService.deduct(null, BigDecimal.valueOf(150));
	}
	
	@Test(expected = InvalidCardOperation.class)
	public void deduct_negativeAmount_throwInvalidCardOperation() throws Exception {
		BigDecimal deductAmount = BigDecimal.valueOf(-200);
		when(cardValidator.validateAmount(deductAmount, CardOperationType.DEDUCT_BALANCE))
			.thenReturn(new InvalidCardOperation("Cannot deduct with negative amount!", CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION));
		cardCommandService.deduct(CARD_SN, deductAmount);
	}

	@Test(expected = InvalidCardOperation.class)
	public void deduct_zeroAmount_throwInvalidCardOperation() throws Exception {
		when(cardValidator.validateAmount(BigDecimal.ZERO, CardOperationType.DEDUCT_BALANCE))
			.thenReturn(new InvalidCardOperation("No amount to deduct!", CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION));
		cardCommandService.deduct(CARD_SN, BigDecimal.ZERO);
	}

	@Test(expected = CardNotFoundException.class)
	public void deduct_cardNotRegistered_throwCardNotFoundException() throws Exception {
		when(cardValidator.validateSerialNumber(CARD_SN)).thenReturn(new CardNotFoundException("Unregistered card!"));
		cardCommandService.deduct(CARD_SN, BigDecimal.valueOf(30));
	}
		
	@Test
	public void deduct_everythingOK_balanceShouldBeCorrectlyDeducted() throws Exception {
		Card card = new Card();
		card.setSerialNo(CARD_SN);
		card.setBalance(BigDecimal.valueOf(500));
		
		when(cardValidator.validateSerialNumber(CARD_SN)).thenReturn(null);
		when(cardRepository.findBySerialNo(CARD_SN)).thenReturn(card);
		
		BigDecimal balance = cardCommandService.deduct(CARD_SN, BigDecimal.valueOf(200));
		assertTrue(balance.compareTo(BigDecimal.valueOf(300)) == 0);
	}
	
	@Test(expected = CardNotFoundException.class)
	public void add_noSerialNumber_throwCardNotFoundException() throws Exception {
		when(cardValidator.validateSerialNumber(null)).thenReturn(new CardNotFoundException("Cannot find card!"));
		cardCommandService.add(null, BigDecimal.valueOf(200));
	}

	@Test(expected = InvalidCardOperation.class)
	public void add_negativeAmount_throwInvalidCardOperation() throws Exception {
		BigDecimal addAmount = BigDecimal.valueOf(-100);
		when(cardValidator.validateAmount(addAmount, CardOperationType.ADD_BALANCE))
			.thenReturn(new InvalidCardOperation("Cannot add with negative amount!", CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION));		
		cardCommandService.add(CARD_SN, addAmount);
	}

	@Test(expected = InvalidCardOperation.class)
	public void add_zeroAmount_throwInvalidCardOperation() throws Exception {
		when(cardValidator.validateAmount(BigDecimal.ZERO, CardOperationType.ADD_BALANCE))
		.thenReturn(new InvalidCardOperation("No amount to add!", CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION));		
		cardCommandService.add(CARD_SN, BigDecimal.ZERO);
	}

	@Test
	public void add_everythingOK_balanceShouldBeCorrectlyAdded() throws Exception {
		Card card = new Card();
		card.setSerialNo(CARD_SN);
		card.setBalance(BigDecimal.ZERO);
		when(cardRepository.findBySerialNo(CARD_SN)).thenReturn(card);
		
		BigDecimal balance = cardCommandService.add(CARD_SN, BigDecimal.valueOf(150));
		assertTrue(balance.compareTo(BigDecimal.valueOf(150)) == 0);
	}
}
package com.yauritux.service.validator.impl;

import java.math.BigDecimal;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.yauritux.exception.CardException;
import com.yauritux.exception.CardNotFoundException;
import com.yauritux.exception.InvalidCardOperation;
import com.yauritux.model.constant.CardExceptionType;
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
@Service
public class CardValidatorImpl implements CardValidator {

	private CardQueryService cardQueryService;

	@Autowired
	CardValidatorImpl(CardQueryService cardQueryService) {
		this.cardQueryService = cardQueryService;
	}
	
	public CardException validateOwner(String owner) {
		if (owner == null || owner.trim().length() == 0) {
			return new InvalidCardOperation("Owner name is required!", CardExceptionType.CARD_MISSING_ATTRIBUTE_EXCEPTION);
		}
		
		return null;
	}

	@Override
	public CardException validateSerialNumber(final String serialNo) {
		if (serialNo == null || serialNo.trim().length() == 0) {
			return new CardNotFoundException("No serial number found !");
		}

		if (!cardQueryService.findBySerialNo(serialNo).isPresent()) {
			return new CardNotFoundException("Cannot find card with serial number '" + serialNo + "' !");
		}

		return null;
	}

	@Override
	public CardException validateAmount(BigDecimal amount, final CardOperationType cardOperationType) {
		switch (cardOperationType) {
		case ADD_BALANCE:
			if (amount.compareTo(BigDecimal.ZERO) < 0) {
				return new InvalidCardOperation("Cannot add balance with negative amount!",
						CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION);
			}
			if (amount.compareTo(BigDecimal.ZERO) == 0) {
				return new InvalidCardOperation("No amount to be added!", CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION);
			}
			break;
		case DEDUCT_BALANCE:
			if (amount.compareTo(BigDecimal.ZERO) < 0) {
				return new InvalidCardOperation("Cannot deduct balance with negative amount!",
						CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION);
			}
			if (amount.compareTo(BigDecimal.ZERO) == 0) {
				return new InvalidCardOperation("No amount to deduct!",
						CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION);
			}
			break;
		case CARD_INITIALIZATION:
			if (amount.compareTo(BigDecimal.ZERO) < 0) {
				return new InvalidCardOperation("Cannot initialize card with negative amount!",
						CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION);
			}
			if (amount.compareTo(BigDecimal.ZERO) == 0) {
				return new InvalidCardOperation("Cannot initialize card with no amount!",
						CardExceptionType.CARD_INVALID_AMOUNT_EXCEPTION);
			}
			break;
		default:
		}

		return null;
	}
}
